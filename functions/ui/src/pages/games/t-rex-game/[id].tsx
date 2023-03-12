import { useParams } from "react-router-dom";
import { TRexGame } from "../../../components/games/TRexGame";
import NoPage from "../../NoPage";

export const TRexGamePage: React.FC = () => {
  const modes = ["default", "corona", "gmp"];
  // get game id from router
  const { id } = useParams();

  if (modes.includes(id)) {
    return <TRexGame mode={id}></TRexGame>;
  } else {
    return <NoPage></NoPage>;
  }
};
