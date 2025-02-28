import styles from "./TimesheetEditor.module.scss";
import { useEffect, useState } from "react";
import TimesheetTable from "../../components/TimesheetTable/TimesheetTable";

function TimesheetEditor({ tableTitle, onSave }) {
  const [tableData, setTableData] = useState([]);

  const onTableUpdate = (data) => {
    console.log(data);
    setTableData(data);
  };

  const onTableSave = () => {
    console.log(tableData);
    window.alert("Timesheet Saved");
    onSave(tableData);
  };

  useEffect(() => {
    console.log("updated table data");
  }, [tableData]);

  return (
    <div className={styles.TimesheetEditor}>
      <div className={styles.TimesheetEditor_Header}>
        <h2>Week Ending: {tableTitle}</h2>
      </div>
      <TimesheetTable
        tableData={tableData}
        onUpdate={onTableUpdate}
      ></TimesheetTable>
      <div className={styles.TimesheetEditor_Footer}>
        <button onClick={onTableSave}>Save Week</button>
      </div>
    </div>
  );
}

export default TimesheetEditor;
