import React from 'react';
import { connect } from "react-redux";
// import { asyncSetTestValue, setTestValue } from "../../redux/actions/testActions";
import CourseGridItem from '../components/CourseGridItem';
import { setTestValue, asyncSetTestValue } from '../../../redux/actions/testActions';
import {Grid, GridRow} from 'semantic-ui-react';
class CourseContainer extends React.Component {
    render() {
        let rows=[]
        for (let i = 0; i <20; i++) {
            rows.push(<Grid.Column key={i} style={{margin: '0px', marginBottom: '25px'}}> <CourseGridItem {...this.props} /></Grid.Column>)
        }
        return (
                <Grid stackable style={{margin: '0px'}}>
                    <GridRow columns={1} only="mobile">{rows}</GridRow>
                    <GridRow columns={2} only="tablet">{rows}</GridRow>
                    <GridRow columns={4} only="computer">{rows}</GridRow>
                </Grid>
                
        );
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

export default connect(mapStateToProps, mapDispatchToProps)(CourseContainer);