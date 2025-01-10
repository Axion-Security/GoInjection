"use client";

import {
    Button, Chip,
    Divider,
    Input,
    Table,
    TableBody,
    TableCell,
    TableColumn,
    TableHeader,
    TableRow
} from "@nextui-org/react";
import {useEffect, useState} from "react";
import {GetTarget, SetTarget, CheckWAF, Interpreter} from "../../../wailsjs/go/main/App";
import {Icon} from "@iconify/react";

export default function Page() {
    const [value, setValue] = useState("");
    const [result, setResult] = useState([]);

    useEffect(() => {
        GetTarget("url").then((result) => {
                    setValue(result);
        });
    }, []);

    async function interpreter() {
        await Interpreter().then((result) => {
            setResult(result);
        }, (err) => {
            console.error(err);
        });
    }

    return (
        <div className={`flex flex-col justify-center items-center p-1`}>
            <div className={`flex flex-row w-full items-center`}>
                <Input
                    size={"md"}
                    placeholder={"Target"}
                    value={value}
                    isReadOnly={true}
                />
                <Button variant={"faded"} className={`ml-2`} onPress={async () => await interpreter()}>Run</Button>
            </div>
            <Divider className={`my-4 w-[700px]`}/>
            {result.length !== 0 ?
                <Table fullWidth={true} w>
                    <TableHeader>
                        <TableColumn>DBMS</TableColumn>
                    </TableHeader>
                    <TableBody>
                        <TableRow key="1">
                            <TableCell>{result}</TableCell>
                        </TableRow>
                    </TableBody>
                </Table>
                :
                <p className="text-zinc-500 uppercase italic tracking-widest">
                    Waiting...
                </p>
            }
            <footer
                className="bg-[#2c2d31] rounded p-2 flex flex-col justify-center items-center mt-4 text-sm text-gray-500 absolute bottom-5 w-[735px]">
                <p>The interpreter determines the SQL language to use when performing SQL injection</p>
            </footer>
        </div>
    )
}