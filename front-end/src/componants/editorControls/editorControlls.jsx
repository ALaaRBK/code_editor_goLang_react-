import React, { Component } from 'react'
import { Button, ListGroup, ListGroupItem, DropdownButton, InputGroup } from 'react-bootstrap'
import DropdownItem from 'react-bootstrap/DropdownItem'
import './editorControls.css'
const languages = ['cpp', 'python', 'go']
const editorControls = (props) => {
    return <div className="editor-controls">

        <Button onclick={evt => {
            var Disabled = document.getElementsByClassName("std")[0].style.display
            if(Disabled == "none"){
                document.getElementsByClassName("std")[0].style.display="block"
            }else{
                document.getElementsByClassName("std")[0].style.display="none"
            }
            document.getElementsByClassName("std")[0] = !Disabled;
        }} className="editor-controls__btn" variant="secondary">
            Stdin
        </Button>

        <DropdownButton className="editor-controls__btn" as={InputGroup.Prepend} variant="secondary" title="Lang">
            {
                languages.map(value => {
                    return <DropdownItem onClick={props.languageClicked} data-lang={value}>
                        {value}
                    </DropdownItem>
                })
            }
        </DropdownButton>
        <Button variant="primary" className="editor-controls__btn editor-controls__btn--primary">
            Run
        </Button>
    </div>
}
export default editorControls