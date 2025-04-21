import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import styles from "./App.module.scss";
import TimesheetApp from "./apps/timesheets/containers/App/App";
import DataContainer from "./containers/DataContainer/DataContainer";

function App() {
  return (
    <div className={styles.App}>
      <DataContainer></DataContainer>
      {/* <TimesheetApp>
      </TimesheetApp> */}
    </div>
  );
}

export default App;
