import { TRexGame } from "../components/TRexGame";
import { TRexGameMode } from "../types/game";

export default function Game() {
  return <TRexGame mode={TRexGameMode.GMP}></TRexGame>
}
