import { Route, Switch } from "react-router-dom";
import React from "react";
import TestContainer from "./containers/TestContainer";
import CourseContainer from "./courses/containers/CourseContainer";

const Routes = () => {
    return (
        <Switch>
            <Route exact path="/" render={() => <div>Welcome</div>}/>
            <Route exact path="/test" component={TestContainer}/>
            <Route exact path="/courses" component={CourseContainer}/>
        </Switch>
    )
}

export default Routes;