import { BrowserRouter, Route, Routes } from "react-router-dom"
import { SignUp } from "./page/SignUp"
import { SignIn } from "./page/SignIn"
import { Landing } from "./page/Landing"
import { Dashboard } from "./page/Dashboard"

function App() {

  return<><BrowserRouter>
  <Routes>
    <Route path="/signup" element={<SignUp />}/>
    <Route path="/signin" element={<SignIn />}/>
    <Route path="/" element={<Landing />}/>
    <Route path="/dashboard" element={<Dashboard />}/>
  </Routes>   
  </BrowserRouter> 
  </>
}

export default App
