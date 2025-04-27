import { lazy, Suspense } from "react";

const RemoteAsset = lazy(() => import("remote_app_1/Asset"));

export const Asset: React.FC = () => {
  return (
    <Suspense fallback="loading">
      <RemoteAsset></RemoteAsset>
    </Suspense>
  );
};
