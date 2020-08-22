import { Link } from "react-router-dom";
import React from "react";

const NavigationLinks = () => {
    return (
        <nav>
            <ul>
                <li>
                    <Link to="/">Dashboard</Link>
                </li>
                <li>
                    <Link to="/courses">Courses</Link>
                </li>
                <li>
                    <Link to="/learning-paths">Learning Path</Link>
                </li>
                <li>
                    <Link to="/events">Events</Link>
                </li>
                <li>
                    <Link to="/events">Events</Link>
                </li>
                <li>
                    <Link to="/setting">Setting</Link>
                </li>
                <li>
                    <Link to="/test">Test</Link>
                </li>
            </ul>
        </nav>
    )
}

export default NavigationLinks;