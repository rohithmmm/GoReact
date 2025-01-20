import { useState } from "react";
import { Header } from "../components/Header";
import axios from "axios";
import { BASE_URL } from "../utils/config";
import { useNavigate } from "react-router-dom";

export function Dashboard(){
    const [clockIn, setClockIn] = useState<string | null>(null);
    const [clockOut, setClockOut] = useState<string | null>(null);
    const redirect = useNavigate();
    
    const token = localStorage.getItem("token");
    const email = localStorage.getItem("email");
    
    function checkToken(){
        if(!token) {
            console.log("Please SignIn again");
            redirect('/signin');
        }
    }
    

    async function handleClockIn(){
        checkToken();
        try {
            const response = await axios.post(BASE_URL + "/clockIn", {
                token,
                email
            })
            console.log(response.data);
            setClockIn(response.data)
        } catch(err) {
            console.error(err);
        }
    }

    async function handleClockOut(){
        checkToken();

        try {
            const response = await axios.post(BASE_URL + "/clockOut", {
                token,
                email
            })
            console.log(response.data);
            setClockOut(response.data)
        } catch(err) {
            console.error(err);
        }
    }

    return<div className="bg-black h-screen flex justify-center">
        <div className="flex flex-col justify-center">
            <Header />
            <div className="text-white">
                <div className="p-2 font-bold text-3xl flex justify-center">
                    Your Timesheet
                </div>
                <div className="p-2 m-6 rounded-md bg-white text-black flex flex-col items-center">
                    <div className="p-2 text-2xl font-semibold">
                        Clock In/Out
                    </div>
                    <div className="mx-10 my-5 flex justify-between font-medium ">
                        <div className="p-6">
                            <button className="px-5 py-3 bg-blue-600 text-white rounded-md" onClick={handleClockIn}>
                                Clock In
                            </button>
                        </div>
                        <div className="p-6">
                            <button className="px-4 py-3 bg-blue-600 text-white rounded-md" onClick={handleClockOut}>
                                Clock Out
                            </button>
                        </div>
                        <div className="text-black font-medium">
                            {clockIn}
                            {clockOut}
                        </div>
                    </div>
                </div>
            </div>
        </div>

    </div>
}