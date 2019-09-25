import React, { Component } from 'react'
import MonacoEditor from 'react-monaco-editor';
class NewCode extends Component {
    constructor(props) {
        super(props)
        this.state = { language: "python" , codeText: "" }
        
    }
    editorDidMount(editor, monaco) {
        // console.log('editorDidMount', editor);
        editor.focus();
    }
    onChange(newValue, e) {
        this.setState({codeText: newValue})
        // console.log('onChange', newValue, e);
    }
    render() {
        return <MonacoEditor
            width="800"
            height="600"
            language={this.state.language}
            theme="vs-dark"
            value={this.state.codeText}
            // options={this}
            onChange={this.onChange}
            editorDidMount={this.editorDidMount}
        />
    }
}
export default NewCode