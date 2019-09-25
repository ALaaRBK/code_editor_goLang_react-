import React, { Component } from 'react'
import Editor, { ControlledEditor } from '@monaco-editor/react';
import EditorControls from '../componants/editorControls/editorControlls';
import Stdin from "../componants/stdin/stdin"
class NewCode extends Component {
    constructor(props) {
        super(props)
        this.state = { language: "python", codeText: "" }

    }
    onChange = (newValue, e) => {
        console.log('onChange', newValue, e);
        this.setState({ codeText: newValue })
    }
    languageClickedHandler = (evt) => {
        this.setState({ language: evt.target.getAttribute("data-lang") })
    }
    render() {


        return (
            <div className="editor-contrainer">
                
                <ControlledEditor onChange={this.onChange} height="500px" width="700px" language={this.state.language} theme="dark" />
                <Stdin />
                <EditorControls languageClicked={this.languageClickedHandler} />
            </div>
        )
    }
}
export default NewCode