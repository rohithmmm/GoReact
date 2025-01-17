import { useState } from "react";
import { Header } from "../components/Header";

export function Dashboard(){
    const [clockIn, setClockIn] = useState<string | null>(null);
    const [clockOut, setClockOut] = useState<string | null>(null);
    

    function handleClockIn(){
        // setClockIn();
    }

    function handleClockOut(){
        // setClockOut();
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