import { useState } from "react";
import styles from "./App.module.scss";
import Login from "../Login/Login";
import Home from "../Home/Home";

function TimesheetApp(props) {
  const [isLoggedIn, login] = useState(false);
  return (
    <div className={styles.App}>
      <Home></Home>
    </div>
  );
}

export default TimesheetApp;
