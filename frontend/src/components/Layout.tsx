import React, { Children, Fragment } from "react"
import MainNavigation from "./MainNavigation"
import classes from './Layout.module.css'
interface Props {
    children: React.ReactNode;
}
const Layout: React.FC<Props> = ({ children }) => {
    return (
        <Fragment>
            <MainNavigation></MainNavigation>
            <main className={classes.main}>{children}</main>
        </Fragment>
    );
};
export default Layout;