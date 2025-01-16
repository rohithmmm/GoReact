import { useNavigate } from "react-router-dom";
import { Icon } from "../assets/icon";

export function Landing(){
    const redirect = useNavigate();
    return<div className="p-2 h-screen bg-black text-white">
        <div className="m-6 flex items-center justify-between">
            <div className="flex items-center">
                <div className="p-2"><Icon /></div>
                <div className="text-3xl font-bold">Timesheet</div>
            </div>
            <div className="flex justify-end p-2 space-x-4">
                <button onClick={()=>redirect('/signup')} className="px-5 py-2 font-medium rounded-md bg-slate-700 text-lg">
                    Register
                </button>
                <button onClick={()=>redirect('/signin')} className="px-5 py-2 font-medium rounded-md text-black bg-white text-lg">
                    Login
                </button>
            </div>
        </div>
        
        <div className="flex flex-col items-center flex-grow justify-center translate-y-6">
            <div className="p-2 text-3xl font-bold">Welcome to Timesheet</div>
            <div className="p-2 text-xl font-normal text-slate-400">Track your work hours with ease</div>
        </div>
    </div>
}