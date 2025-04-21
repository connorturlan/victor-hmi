import { useEffect, useState } from "react";
import { usePub } from "../../core/datasources";
import styles from "./DataPub.module.scss";

function DataPub({ controller }) {
  const [time, setTime] = useState(0);

  const publish = usePub();

  const updateTime = () => {
    setTime((prevTime) => prevTime + 1);
  };

  useEffect(() => {
    console.log("useEffect runs");
    const interval = setInterval(updateTime, 1_000);
    return () => clearInterval(interval);
  }, []);

  useEffect(() => {
    publish("onUpdateTime", time);
  }, [time]);

  return <div>publish time {time}</div>;
}

export default DataPub;
