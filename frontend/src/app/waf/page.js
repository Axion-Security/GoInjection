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
import {GetTarget, SetTarget, CheckWAF} from "../../../wailsjs/go/main/App";
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

    async function detectWAF() {
        await CheckWAF().then((result) => {
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
                <Button variant={"faded"} className={`ml-2`} onPress={async () => await detectWAF()}>Run</Button>
            </div>
            <Divider className={`my-4 w-[700px]`}/>
            {result.length !== 0 ?
                <Table fullWidth={true} w>
                    <TableHeader>
                        <TableColumn>DOMAIN</TableColumn>
                        <TableColumn>IS WAF?</TableColumn>
                        <TableColumn>WAF TYPE/REASON</TableColumn>
                    </TableHeader>
                    <TableBody>
                        <TableRow key="1">
                            <TableCell>{result.domain}</TableCell>
                            <TableCell>{
                             result.isWaf === "True" ?
                                <Chip className="capitalize" color={"success"} size="sm" variant="flat">
                                    True
                                </Chip>
                                : result.isWaf === "Potential" ?
                                <Chip className="capitalize" color={"warning"} size="sm" variant="flat">
                                    Potential
                                </Chip>
                                :
                                <Chip className="capitalize" color={"danger"} size="sm" variant="flat">
                                    False
                                </Chip>
                            }</TableCell>
                            <TableCell>
                                {
                                    result.wafType === "" ?
                                        <p>NONE</p>
                                        :
                                        <p>{result.wafType}</p>
                                }
                            </TableCell>
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
           <p>Checks which WAF the site uses using different techniques</p>
            </footer>
        </div>
    )
}