import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import styles from "./App.module.scss";
import TimesheetApp from "./apps/timesheets/containers/App/App";
import DataContainer from "./containers/DataContainer/DataContainer";
import HostContainer from "./containers/HostContainer/HostContainer";

function App() {
  return (
    <div /* className={styles.App} */>
      <p>mfe should be below</p>
      {/* <HostContainer /> */}
      <p>mfe should be above</p>
      {/* <TimesheetApp>
      </TimesheetApp> */}
    </div>
  );
}

export default App;
