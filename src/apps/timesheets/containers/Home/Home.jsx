import { useEffect, useState } from "react";
import Table from "../../components/Table/Table";
import styles from "./Home.module.scss";
import TimesheetTable from "../../components/TimesheetTable/TimesheetTable";
import TimesheetEditor from "../TimesheetEditor/TimesheetEditor";
import WeekTable from "../../components/WeekTable/WeekTable";

function Home() {
  const [weeks, setWeeks] = useState([]);
  const [working, setWorking] = useState("");

  const handleWeekSelect = (week) => {
    setWorking(week);
  };

  const handleSave = (week) => {
    setWorking("");
  };

  const child = !working ? (
    <WeekTable
      weeks={weeks}
      onClick={handleWeekSelect}
      onUpdate={setWeeks}
    ></WeekTable>
  ) : (
    <TimesheetEditor
      tableTitle={"sum fkn name here"}
      onSave={handleSave}
    ></TimesheetEditor>
  );

  return (
    <div className={styles.Home}>
      <div className={styles.Home_Header}>
        <h1>Timesheets</h1>
      </div>
      {child}
      <div className={styles.Home_Footer}></div>
    </div>
  );
}

export default Home;
