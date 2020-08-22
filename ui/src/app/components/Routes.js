import { Route, Switch } from "react-router-dom";
import React from "react";
import TestContainer from "../containers/TestContainer";

const Routes = () => {
    return (
        <Switch>
            <Route exact path="/" render={() => <div>Welcome</div>}/>
            <Route exact path="/test" component={TestContainer}/>
        </Switch>
    )
}

export default Routes;