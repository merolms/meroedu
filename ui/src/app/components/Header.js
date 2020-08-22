import React from 'react';
import { Button, Menu } from "semantic-ui-react";

const Header = (props) => {
    return (
        <Menu secondary className="header bg-white">
            <Menu.Item>
                <a className="primary h1" href="#home">Mero Edu</a>
            </Menu.Item>
            <Menu.Menu position='right'>
            <Menu.Item>
                <Button icon="bell" className="bg-white"/>
            </Menu.Item>
            <Menu.Item>
                <Button icon="ellipsis vertical" className="bg-white"/>
            </Menu.Item>
        </Menu.Menu>
      </Menu>
    )
}

export default Header;
