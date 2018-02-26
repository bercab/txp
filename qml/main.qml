import QtQuick 2.7
import QtQuick.Controls 2.1
import QtQuick.Dialogs 1.2
import QtQuick.Layouts 1.0


ApplicationWindow {
  id: window
  visible: true
  title: "Conversor TXP SEPA"
  minimumWidth: 400
  minimumHeight: 400
  background: Image {
    source: "cnp.jpeg"
  }
  footer: ToolBar {
      RowLayout {
          anchors.fill: parent
          Label { 
            text: "TXP v0.1 · Copyright © 2017 FerriolSoft" 
            //elide: Label.ElideRight
            horizontalAlignment: Qt.AlignHCenter
            verticalAlignment: Qt.AlignVCenter
            Layout.fillWidth: true
          }
      }
  }
  Connections {
    target: QmlBridge
    onParsedFile: {
      outtext.text = data
    }
    onInfoMessage: {
      messageDialog.text = data
      messageDialog.title = "Info"
      messageDialog.icon = StandardIcon.Information
      messageDialog.open()
    }
    onErrorMessage: {
      messageDialog.text = data
      messageDialog.title = "Error"
      messageDialog.icon = StandardIcon.Critical
      messageDialog.open()
    }
    onSuccessSave: {
      successDialog.text = data
      successDialog.open()
    }
  }

  Column {
    anchors.centerIn: parent
    spacing: 5
    //anchors.fill: parent
    // Label { 
    //   text: "Resultados:" 
    // }
    TextArea {
      id: outtext
          background: Rectangle {
          border.color: "#666666" 
    }
      anchors.horizontalCenter: parent.horizontalCenter
      width: 330
      height: 200 
      text: qsTr("Text Area")
      readOnly: true
      visible: false
    }

    Button {
      id: openButton
      anchors.horizontalCenter: parent.horizontalCenter
      text: "Selecciona fichero TXP"
      onClicked: {
        openDialog.open()
      }
    }
    Button {
      id: saveButton
      anchors.horizontalCenter: parent.horizontalCenter
      text: "Guardar fichero XML"
      onClicked: saveDialog.open()
      visible: false
    }

  }


  FileDialog {
    id: openDialog
    title: "Por favor elige fichero txt"
    folder: shortcuts.home
    nameFilters: [ "Ficheros TXP (*.txt)", "Todos los ficheros (*)" ]
    selectedNameFilter: "Ficheros TXP (*.txt)"
    onAccepted: {
      console.log("hola - " + openDialog.fileUrl)
      QmlBridge.acceptedOpenFile(openDialog.fileUrl)
      outtext.visible = true
      openButton.visible = false
      saveButton.visible = true
      saveDialog.folder = openDialog.folder
    }
    onRejected: {
      console.log("Cancelado")
    }
    // Component.onCompleted: visible=true
  }
  FileDialog {
    id: saveDialog
    title: "Guardar XML"
    selectExisting: false
    folder: shortcuts.home
    nameFilters: [ "Ficheros XML (*.xml)", "Todos los ficheros (*)" ]
    selectedNameFilter: "Ficheros XML (*.xml)"
    onAccepted: {
      console.log("hola - " + saveDialog.fileUrl)
      QmlBridge.acceptedSaveFile(saveDialog.fileUrl)
      outtext.visible = true
      openButton.visible = false
    }
    onRejected: {
      console.log("Cancelado")
    }
    // Component.onCompleted: visible=true
  }

  MessageDialog {
    id: messageDialog
    title: "Información"
    text: ""
    visible: false
    icon: StandardIcon.Information
    //Component.onCompleted: visible = true
  }
  MessageDialog {
    id: successDialog
    title: "Proceso Finalizado OK. Viva la Ferriol!"
    text: ""
    visible: false
    icon: StandardIcon.Information
    onAccepted:  Qt.quit()
  }
}