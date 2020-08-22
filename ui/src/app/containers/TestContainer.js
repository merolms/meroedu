import React from 'react';
import { connect } from "react-redux";
import TestComponent from "../components/TestComponent";
import { asyncSetTestValue, setTestValue } from "../../redux/actions/testActions";

class TestContainer extends React.Component {
    render() {
        return <TestComponent {...this.props} />;
    }
}

const mapStateToProps = state => {
    return {
        test: state.test
    }
};

const mapDispatchToProps = dispatch => {
    return {
        setValue: (value) => {
            dispatch(setTestValue(value))
        },
        asyncSetValue: (value) => {
            dispatch(asyncSetTestValue(value));
        }
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(TestContainer);