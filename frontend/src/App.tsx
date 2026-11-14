import { useState } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";
import { Create } from "../wailsjs/go/notes/Service";

function App() {
  const [resultText, setResultText] = useState(
    "Please enter your name below 👇",
  );
  const [name, setName] = useState("");
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  function createNote() {
    Create("tchau paizao de la massa");
  }
  return (
    <div id="App">
      <button onClick={createNote}>criar nota</button>
    </div>
  );
}

export default App;
