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


export function SignUp(){
    const [isLoading, setIsLoading] = useState(false);

    const firstNameRef = useRef<HTMLInputElement>(null);
    const lastNameRef = useRef<HTMLInputElement>(null);
    const emailRef = useRef<HTMLInputElement>(null);
    const passwordRef = useRef<HTMLInputElement>(null);
    const redirect = useNavigate();

    async function handleClick(){
        
        const first_name = firstNameRef.current?.value;
        const last_name = lastNameRef.current?.value;
        const email = emailRef.current?.value; 
        const password = passwordRef.current?.value; 

        console.log("email: ",email);
        console.log("password: ",password);
        console.log("firstName: ",first_name);
        console.log("lastName: ",last_name);

        if(!first_name || !last_name || !email || !password) {
            console.log("please fill all the details");
            return;
        }

        setIsLoading(true);

        try {
            const response = await axios.post(BASE_URL + "/addEmployee",{
                first_name,
                last_name,
                email,
                password
            })
            console.log(response.data);
            redirect('/signin');
        } catch(error) {
            console.error(error);
            setIsLoading(false);
        }
    }

    return<div className="bg-black h-screen flex justify-center items-center">
        <div className="flex flex-col justify-center">
        <Header/>
            <div className="p-2 px-4 w-80 text-center rounded-lg border border-slate-800 shadow-md shadow-slate-800">
                <Heading label="Sign Up"/>
                <SubHeading label={"Please fill the below fields"} />
                <InputBox inputRef={firstNameRef} label={"First name"} placeholder={"John"} />
                <InputBox inputRef={lastNameRef} label={"Last name"} placeholder={"Doe"} />
                <InputBox inputRef={emailRef} label={"Email"} placeholder={"johndoe@gmail.com"} />
                <InputBox inputRef={passwordRef} label={"Password"} placeholder={"12345"} />
                <div className="pt-4">
                    <Button label={isLoading ? "Loading..." : "Sign Up"} onClick={handleClick} disabled={isLoading}/>
                </div>
                <Warning label={"Already Registered?"} linkText={"Sign in"} to={"/signin"} />
            </div>
        </div>
    </div>
}