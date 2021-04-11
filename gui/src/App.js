import React from 'react';
import './App.css';
import Navbar from "./component/Navbar";
import Storm from "./component/Storm";
import BitCoin from "./component/BitCoin";
import {
    BrowserRouter as Router,
    Switch,
    Route
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
                    <Route path="/btc" component={BitCoin}>
                        <BitCoin />
                    </Route>
                </Switch>
            </Router>
        </div>
    )
}

export default App;
