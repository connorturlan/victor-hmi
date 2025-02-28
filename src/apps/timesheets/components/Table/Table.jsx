import { useEffect, useState } from "react";
import styles from "./Table.module.scss";

function Table({ headings, tableData, onUpdate, onClick }) {
  // const [tableData, setRows] = useState(data || []);

  const updateRow = (rowIndex, cellIndex, value) => {
    const newTable = tableData.slice();
    newTable[rowIndex][cellIndex] = value;
    onUpdate && onUpdate(newTable);
  };

  const addRow = () => {
    const newTable = tableData.slice();
    const newArray = Array.apply(0, Array(headings.length)).map(() => {
      return -1;
    });
    newArray[0] = "";
    newTable.push(newArray);
    onUpdate && onUpdate(newTable);
  };

  const removeRow = (rowIndex) => {
    console.log(rowIndex);
    const preTable = tableData.slice(0, rowIndex);
    const postTable = tableData.slice(rowIndex + 1);
    const newTable = preTable.concat(postTable);
    console.log(preTable, postTable, newTable);
    onUpdate && onUpdate(newTable);
  };

  // useEffect(() => {
  //   console.log("updating table");
  //   onUpdate && onUpdate(tableData);
  // }, [tableData]);

  // useEffect(() => {
  //   console.log("table was updated");
  // }, [data]);

  return (
    <div className={styles.Table}>
      <table className={styles.Table_Table}>
        <thead>
          <tr>
            {headings.map((heading) => {
              return <th key={`heading${heading}`}>{heading}</th>;
            })}
          </tr>
        </thead>
        <tbody>
          {tableData &&
            tableData.map((row, rowIndex) => {
              return (
                <tr key={`row${rowIndex}`} onClick={onClick}>
                  {row.slice(0, 1).map((cell, cellIndex) => {
                    return (
                      <td key={`cell${rowIndex},${cellIndex}`}>
                        <input
                          placeholder="TIMECODE"
                          value={cell}
                          onChange={(event) =>
                            updateRow(rowIndex, cellIndex, event.target.value)
                          }
                        ></input>
                      </td>
                    );
                  })}
                  {row.slice(1).map((cell, cellIndex) => {
                    return cell < 0 ? (
                      <td></td>
                    ) : (
                      <td key={`cell${rowIndex},${cellIndex}`}>
                        <input
                          placeholder="0"
                          value={cell}
                          onChange={(event) =>
                            updateRow(
                              rowIndex,
                              cellIndex + 1,
                              event.target.value
                            )
                          }
                        ></input>
                      </td>
                    );
                  })}
                  <td>
                    <button onClick={() => removeRow(rowIndex)}>
                      Remove Row
                    </button>
                  </td>
                </tr>
              );
            })}
          <tr>
            <td>
              <button onClick={addRow}>Add Row</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}

export default Table;
