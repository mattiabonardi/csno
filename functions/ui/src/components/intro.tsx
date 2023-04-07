import "../styles/components/intro.css";
import { Loading } from "./loading";

interface IntroProps {
  active: boolean;
}

export const Intro: React.FC<IntroProps> = ({ active }) => {
  if (active) {
    return (
      <>
        <div className="wrapper">
          <Loading></Loading>
        </div>
      </>
    );
  } else {
    return <></>;
  }
};
