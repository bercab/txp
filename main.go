package main

import (
	"fmt"
	"log"
	"os"

	"github.com/apsl/sepakit/convert"
	"github.com/apsl/sepakit/sepadebit"
	"github.com/pkg/errors"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quick"
)

var docxml *sepadebit.Document

func main() {
	//os.Setenv("QMLSCENE_DEVICE", "softwarecontext")
	//os.Setenv("QT_QUICK_BACKEND", "software")
	quick.QQuickWindow_SetSceneGraphBackend2("software")
	app := gui.NewQGuiApplication(len(os.Args), os.Args)

	// Enable high DPI scaling
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Use the material style for qml
	// quickcontrols2.QQuickStyle_SetStyle("material")
	var qmlBridge = NewQmlBridge(nil)
	var err error

	qmlBridge.ConnectAcceptedOpenFile(func(data string) {
		docxml, err = parseTxt(data)
		if err != nil {
			qmlBridge.ErrorMessage(fmt.Sprintf("Error: %s", err))
			return
		}
		qmlBridge.ParsedFile(getResumeData(docxml))
		return
	})
	qmlBridge.ConnectAcceptedSaveFile(func(data string) {
		fmt.Println("saved: ", data)
		path, err := writeXML(data)
		if err != nil {
			qmlBridge.ErrorMessage(fmt.Sprintf("%s\n", err))
			return
		}
		qmlBridge.SuccessSave(fmt.Sprintf("XML Grabado correctamente en %s\n", path))
		return
	})
	// Create a QML application engine
	engine := qml.NewQQmlApplicationEngine(nil)
	engine.RootContext().SetContextProperty("QmlBridge", qmlBridge)

	// Load the main qml file
	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	// Execute app
	gui.QGuiApplication_Exec()
}

func writeXML(url string) (string, error) {
	path := getPath(url)
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return path, errors.Wrap(err, fmt.Sprintf("Error abriendo %s", path))
	}
	err = docxml.WriteLatin1(f)
	if err != nil {
		return path, errors.Wrap(err, fmt.Sprintf("Error grabando %s", path))
	}
	return path, nil
}

//parseTxt opens a file from a given file:// path and returns doctxt.
func parseTxt(url string) (*sepadebit.Document, error) {
	// path = strings.TrimPrefix(path, "file://")
	path := getPath(url)
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, fmt.Sprintf("Error opening %s", path))
	}
	defer f.Close()
	docxml, err = convert.Latin1DebitTxtToXMLDoc(f)
	if err != nil {
		log.Printf("Parsing error: %s", err)
		return nil, errors.Wrap(err, "Parsing error")
	}
	return docxml, nil
}

func getPath(url string) string {
	qurl := core.NewQUrl3(url, core.QUrl__TolerantMode)
	path := qurl.ToLocalFile()
	return path
}

func getResumeData(docxml *sepadebit.Document) string {
	resume := fmt.Sprintf("Resumen Datos XML:\n\n")
	resume = resume + fmt.Sprintf("Entidad: %s\n", docxml.InitiatingParty.Name)
	resume = resume + fmt.Sprintf("IBAN: %s\n", docxml.Payments[0].IBAN)
	resume = resume + fmt.Sprintf("Fecha efecto: %s\n", docxml.Payments[0].RequestedCollectionDate)
	resume = resume + fmt.Sprintf("Transacciones: %d\n", docxml.TransacNb)
	resume = resume + fmt.Sprintf("Total: %s €\n", docxml.CtrlSum)
	resume = resume + fmt.Sprintf("\nSi es correcto, selecciona Guardar\n")

	return resume
}

type QmlBridge struct {
	core.QObject

	_ func(data string) `signal:"parsedFile"`
	_ func(data string) `signal:"errorMessage"`
	_ func(data string) `signal:"infoMessage"`
	_ func(data string) `signal:"successSave"`
	_ func(data string) `slot:"acceptedOpenFile"`
	_ func(data string) `slot:"acceptedSaveFile"`
}
