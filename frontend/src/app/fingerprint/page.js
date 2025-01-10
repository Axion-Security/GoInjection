"use client";

import {
    Button, Chip,
    Divider,
    Input, Link, Select, SelectItem,
    Table,
    TableBody,
    TableCell,
    TableColumn,
    TableHeader,
    TableRow
} from "@nextui-org/react";
import {useEffect, useState} from "react";
import {GetTarget, SetTarget, CheckWAF, Fingerprint} from "../../../wailsjs/go/main/App";
import {Icon} from "@iconify/react";

const methods = [
    {key: "union", label: "Union-based"},
    {key: "error", label: "Error-based"},
    {key: "stacked", label: "Stacked-based"},
];

export default function Page() {
    const [method, setMethod] = useState(new Set([""]));
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

    async function fingerprintDB() {
        await Fingerprint(method.currentKey).then((result) => {
            setResult(result);
        }, (err) => {
            console.error(err);
        });
    }

    return (
        <div className={`flex flex-col justify-center items-center p-1`}>
            <div className={`flex flex-row w-full items-center gap-2`}>
                <Input
                    size={"sm"}
                    label="Target"
                    placeholder="Enter Url with Parameters"
                    value={value}
                    onChange={(e) => setTarget(e.target.value)}
                />
                <Select
                    isRequired={true}
                    className="max-w-[140px]"
                    size={"sm"}
                    label="Select Method"
                    selectedKeys={method}
                    onSelectionChange={setMethod}
                >
                    {methods.map((method) => (
                        <SelectItem key={method.key}>{method.label}</SelectItem>
                    ))}
                </Select>
                <Button variant={"faded"} className={`ml-2`} onPress={async () => await fingerprintDB()}>Run</Button>
            </div>
            <Divider className={`my-4 w-[700px]`}/>
            {result.length !== 0 ?
                <Table fullWidth={true} w>
                    <TableHeader>
                        <TableColumn>DOMAIN</TableColumn>
                        <TableColumn>DBMS</TableColumn>
                    </TableHeader>
                    <TableBody>
                        <TableRow key="1">
                            <TableCell>{result.domain}</TableCell>
                            <TableCell>{result.dbms}</TableCell>
                        </TableRow>
                    </TableBody>
                </Table>
                :
                <p className="text-zinc-500 uppercase italic tracking-widest">
                    Waiting...
                </p>
            }
            <p className={`text-sm text-white/50 mt-2 line-through`}>
                <b className={`mr-1 text-red-500`}>
                    *
                </b>
                Stacked and Union-based methods require columns to be specified. <Link href={"/resolver"}
                                                                                       showAnchorIcon={true}
                                                                                       size={"sm"}>Specify
                Columns</Link>
            </p>
            <footer
                className="bg-[#2c2d31] rounded p-2 flex flex-col justify-center items-center mt-4 text-sm text-gray-500 absolute bottom-5 w-[735px]">
                <p>Determines the DBMS using the selected method</p>
            </footer>
        </div>
    )
}