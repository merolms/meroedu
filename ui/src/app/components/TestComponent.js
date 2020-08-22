import React from 'react';

const TestComponent = (props) => {
    return (
        <div>
            <div>
                <button onClick={() => {props.setValue(!props.test.ok)}}>Sync Toggle</button>
                <button onClick={() => {props.asyncSetValue(!props.test.ok)}}>Async Toggle</button>
            </div>
            <pre>{JSON.stringify(props.test)}</pre>
        </div>
    )
}

export default TestComponent;