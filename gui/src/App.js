import React, {useState, useEffect} from 'react';
import './App.css';
import Navbar from "./component/Navbar";
import Storm from "./component/Storm";
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
} from "react-router-dom";

const App = () => {
    return (
        <div className="App">
            <Router>
                <Navbar />
                <Switch>
                    <Route path="/storm" component={Storm}>
                        <Storm />
                    </Route>
                </Switch>
            </Router>
        </div>
    )
}

export default App;
