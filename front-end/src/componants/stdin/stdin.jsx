import  React,{Component} from "react"
import {Form}from "react-bootstrap"
import "./stdin.css"
const stdin = (props)=>{

    return (
        <div className="std">
            <Form>
                <Form.Group className="std-containar">
                    <Form.Label>
                        Enter Input 
                    </Form.Label>
                    <Form.Control as="textarea" className="std-textarea" rows="5"/>
                </Form.Group>
            </Form>
        </div>
    )
}
export default stdin