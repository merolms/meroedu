import React from 'react';
import { Card, Icon, Image } from 'semantic-ui-react'
import styled from 'styled-components';

const ExtraItem = styled.div `
    color: #141313;
    font-weight: bolder;
    margin: 0px 5px;
`

const CourseGridItem = (props) => {
    return (
        <Card fluid={true} style={{width:'328px'}}>
            <Image src='https://react.semantic-ui.com/images/avatar/large/daniel.jpg' wrapped />
            <Card.Content>
                <Card.Header>Graphic Design Work Work Lorem Lipsd UI Lorem Lipsonasdfad</Card.Header>
                {/* <Card.Meta>
                    <span className='date'>Joined in 2015</span>
                </Card.Meta> */}
                <Card.Description>
                    Graphic Design Work Lorem Lipson UI Lorem Lipson UI Lorem Lipson Graphic Design Work Lorem Lipson UInd Lorem Lipson UI Lorem Lipson
                </Card.Description>
            </Card.Content>
            <Card.Content extra style={{display: 'inline-flex'}}>
                <ExtraItem><Icon name='user' />Course Type</ExtraItem>
                <ExtraItem><Icon name='user' />11 Dec 2020</ExtraItem>
                <ExtraItem><Icon name='user' />50</ExtraItem>
            </Card.Content>
    </Card>
    )
}
export default CourseGridItem;
