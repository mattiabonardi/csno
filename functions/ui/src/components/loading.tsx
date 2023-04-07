import { createRef, useEffect } from "react";
import  "../styles/components/loading.css";

export const Loading = () => {
  const bar = createRef<HTMLDivElement>();

  useEffect(() => {
    var width = 1;
    var id = setInterval(loading, 4);
    function loading() {
      if (width >= 100) {
        clearInterval(id);
      } else {
        width++;
        (bar.current as HTMLDivElement).style.width = width + "%";
      }
    }
  }, [bar]);

  return (
    <div className="container">
      <div ref={bar} className="bar"></div>
    </div>
  );
};