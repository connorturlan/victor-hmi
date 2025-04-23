import React, { lazy, Suspense } from "react";
import styles from "./HostContainer.module.scss";
// import {
//   __federation_method_getRemote,
//   __federation_method_setRemote,
//   // @ts-ignore
// } from "module-federation/vite";

function HostContainer(props) {
  const [mfePath, setMfe] = useState("");

  return (
    <Suspense fallback="Loading">
      {/* {mfePath === "" ? (
        <input
          onSubmit={(event) => {
            setMfe(event.target.value);
          }}
        ></input>
      ) : (
        lazy(() => {
          const name = "example-mfe";
          const module = "./example-mfe";

          __federation_method_setRemote(name, {
            url: () => Promise.resolve(mfePath),
            format: "esm",
            from: "vite",
          });

          __federation_method_getRemote(name, module);
        })
      )} */}
    </Suspense>
  );
}

export default HostContainer;
