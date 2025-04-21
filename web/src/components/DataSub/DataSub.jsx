import { useState } from "react";
import { DataSubscriber, UseSub } from "../../core/datasources";
import styles from "./DataSub.module.scss";

const sub = new DataSubscriber();

function DataSub({ controller }) {
  const [time, setTime] = useState(0);

  UseSub("onUpdateTime", (data) => {
    setTime(data);
  });

  return <h2>current time is {time}</h2>;
}

export default DataSub;
