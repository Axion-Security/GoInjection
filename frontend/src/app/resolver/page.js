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
import {GetTarget, SetTarget, CheckWAF, Resolver} from "../../../wailsjs/go/main/App";
import {Icon} from "@iconify/react";

export default function Page() {
    const [value, setValue] = useState("");
    const [result, setResult] = useState([]);

    useEffect(() => {
        async function fetchData() {
            const result = await GetTarget("url");
            setValue(result);
        }
        fetchData();
    }, []);

    function setTarget(value) {
        SetTarget("url", value);
        setValue(value);
    }

    async function resolver() {
        await Resolver().then((result) => {
            setResult(result);
        }, (err) => {
            console.error(err);
        });
    }

    return (
        <div className={`flex flex-col justify-center items-center p-1`}>
            <div className={`flex flex-row w-full items-center`}>
                <Input
                    size={"sm"}
                    label="Target"
                    placeholder="Enter Url with Parameters"
                    value={value}
                    onChange={(e) => setTarget(e.target.value)}
                />
                <Button variant={"faded"} className={`ml-2`} onPress={async () => await resolver()}>Run</Button>
            </div>
            <Divider className={`my-4 w-[700px]`}/>
            {result.length !== 0 ?
                <Table fullWidth={true} w>
                    <TableHeader>
                        <TableColumn>DOMAIN</TableColumn>
                        <TableColumn>COLUMNS</TableColumn>
                        <TableColumn>DATABASE NAME</TableColumn>
                    </TableHeader>
                    <TableBody>
                        <TableRow key="1">
                            <TableCell>{result.domain}</TableCell>
                            <TableCell>{result.columns}</TableCell>
                            <TableCell>{result.database}</TableCell>
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

                <p>Identifies the number of columns and retrieves the database name from the target</p>
            </footer>
        </div>
    )
}