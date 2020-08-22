import { BrowserRouter as Router } from "react-router-dom";
import React from "react";
import Routes from "./Routes";
import NavigationLinks from "./NavigationLinks";
import UserProfileInfo from "./UserProfileInfo";

const Sidebar = (props) => {
	return (
		<Router>
			<div className="sidebar bg-white">
				<UserProfileInfo
					image="https://picsum.photos/100/100"
					primaryText="Angelina Doe"
					secondaryText="Project Manager"
				/>
				<NavigationLinks />
			</div>

			<Routes />
		</Router>
	);
};

export default Sidebar;
