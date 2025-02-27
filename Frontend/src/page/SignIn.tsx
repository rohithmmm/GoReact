import axios from "axios";
import { useRef, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button } from "../components/Button";
import { Header } from "../components/Header";
import { Heading } from "../components/Heading";
import { InputBox } from "../components/InputBox";
import { SubHeading } from "../components/SubHeading";
import { Warning } from "../components/Warning";
import { BASE_URL } from "../utils/config";

export function SignIn(){
    const [isLoading, setIsLoading] = useState(false);
    
    const emailRef = useRef<HTMLInputElement>(null);
    const passwordRef = useRef<HTMLInputElement>(null);
    const redirect = useNavigate();

    async function handleClick(){
        const email = emailRef.current?.value;
        const password = passwordRef.current?.value;
        if(!email || !password) {
            console.log("please fill all the details");
            return;
        }

        setIsLoading(true);

        try {
            const response = await axios.post(BASE_URL + "/signIn",{
                email,
                password
            })
            console.log(response.data);
            localStorage.setItem("token", response.data.token);
            redirect('/dashboard');
        } catch(error) {
            console.error(error);
            setIsLoading(false);
        }
    }

    return<div className="bg-black h-screen flex justify-center">
        <div className="flex flex-col justify-center">
        <Header/>
            <div className="p-2 px-4 w-80 h-max text-center rounded-lg border border-slate-800 shadow-md shadow-slate-800">
                <Heading label="Sign In"/>
                <SubHeading label={"Enter your credentials"} />
                <InputBox inputRef={emailRef} label={"Email"} placeholder={"johndoe@gmail.com"} />
                <InputBox inputRef={passwordRef} label={"Password"} placeholder={"12345"} />
                <div className="pt-4">
                    <Button label={isLoading ? "Loading..." : "Sign In"} onClick={handleClick} disabled={isLoading}/>
                </div>
                <Warning label={"Don't have an account?"} linkText={"Sign up"} to={"/signup"} />
            </div>
        </div>
    </div>
    
}