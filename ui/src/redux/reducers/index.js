import { testReducer } from "./testReducer";
import { combineReducers } from "redux";

export default combineReducers({
    test: testReducer
});