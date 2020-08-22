import { SET_VALUE } from "./types/testActionTypes";
import { getTest } from "../../app/api/testApi";

export const setTestValue = (value) => {
    return {
        type: SET_VALUE,
        payload: value
    };
};

export const asyncSetTestValue = (value) => {
    return async dispatch => {
        await getTest();
        dispatch(setTestValue(value));
    }
}