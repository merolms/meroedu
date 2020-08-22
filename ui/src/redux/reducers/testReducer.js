import { SET_VALUE } from "../actions/types/testActionTypes";

const initialState = {
    ok: true
};

export const testReducer = (state = initialState, action) => {
    switch(action.type) {
        case SET_VALUE:
            return {
                ok: action.payload
            };

        default:
            return state;
    }
};