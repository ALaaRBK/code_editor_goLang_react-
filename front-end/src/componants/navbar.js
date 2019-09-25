import React from "react"
import { Navbar, Nav } from "react-bootstrap"
const NavbarHeader = props => {
    return <Navbar bg="dark" variant="dark">
        <Navbar.Brand href="/">online editor</Navbar.Brand>
        <Nav className="mr-auto">
            <Nav.Link href="/">New Code</Nav.Link>
            <Nav.Link href="/recent">Recent Codes</Nav.Link>
        </Nav>
    </Navbar>
}
export default NavbarHeader