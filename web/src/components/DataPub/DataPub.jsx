import { useEffect, useState } from "react";
import { usePub } from "../../core/datasources";
import styles from "./DataPub.module.scss";

function DataPub({ controller }) {
  const [time, setTime] = useState(0);

  const publish = usePub();

  const updateTime = () => {
    publish("onUpdateTime", time);
    setTime(time + 1);
  };

  useEffect(() => {
    setInterval(updateTime, 1_000);
  }, []);

  return <button onClick={updateTime}>publish time {time}</button>;
}

export default DataPub;
