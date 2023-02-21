import React, {useEffect, useState} from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  const [index, setIndex] = useState([]);
  console.log("running 1");
  useEffect(() => {
    fetch('http://localhost:8000/')
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
          setIndex(data);
        })
        .catch((err) => {
          console.log(err.message);
        });
  }, []);

  return (
    <div className="App">
      {index}
    </div>
  );
}

export default App;
