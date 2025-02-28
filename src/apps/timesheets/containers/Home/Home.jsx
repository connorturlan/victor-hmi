import { useState } from "react";
import Table from "../../components/Table/Table";
import styles from "./Home.module.scss";

function Home(props) {
  const [tableData, setTableData] = useState(null);

  const onTableUpdate = (data) => {
    console.log(data);
    setTableData(data);
  };

  const onTableSave = () => {
    console.log(tableData);
    window.alert("Timesheet Saved");
  };

  return (
    <div className={styles.Home}>
      <Table
        headings={["Timecode", "Monday", "Thursday", "Saturday"]}
        data={[
          [0, 0, 1, 1],
          [1, 0, 1, 2],
          [2, 0, 1, 2],
          [3, 0, 1, 2],
        ]}
        onUpdate={onTableUpdate}
      ></Table>
      <div className={styles.Home_Footer}>
        <button onClick={onTableSave}>Save</button>
      </div>
    </div>
  );
}

export default Home;
