import React from 'react'
import '../App.css';
import * as ReactBootStrap from "react-bootstrap";
import {Link} from 'react-router-dom'

const Navbar = () => {
    return(
        <div className="App">
            <ReactBootStrap.Navbar bg="light" expand="lg">
                <ReactBootStrap.Navbar.Brand href="/">Esquire Follow Up</ReactBootStrap.Navbar.Brand>
                <ReactBootStrap.Navbar.Toggle aria-controls="basic-navbar-nav" />
                <ReactBootStrap.Navbar.Collapse id="basic-navbar-nav">
                <ReactBootStrap.Nav className="mr-auto">
                    <Link to="/storm">
                        <ReactBootStrap.Nav.Link href="#storm">Storm Events</ReactBootStrap.Nav.Link>
                    </Link>
                    <Link to="/btc">
                        <ReactBootStrap.Nav.Link href="#storm">BitCoin</ReactBootStrap.Nav.Link>
                    </Link>
                </ReactBootStrap.Nav>
                </ReactBootStrap.Navbar.Collapse>
            </ReactBootStrap.Navbar>
        </div>
    )
}

export default Navbar;