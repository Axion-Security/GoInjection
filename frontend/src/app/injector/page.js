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
    TableRow, Tooltip
} from "@nextui-org/react";
import {useEffect, useState} from "react";
import {GetTarget, SetTarget, CheckWAF, Injection} from "../../../wailsjs/go/main/App";
import {Icon} from "@iconify/react";
import {ClipboardSetText} from "../../../wailsjs/runtime";

export default function Page() {
    const [value, setValue] = useState("");
    const [unionResult, setUnion] = useState([]);
    const [errorResult, setError] = useState([]);
    const [blindBooleanResult, setBlindBoolean] = useState([]);
    const [blindTimeResult, setBlindTime] = useState([]);

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

    async function inject() {
        const unionResult = await Injection("union");
        setUnion(unionResult);

        const errorResult = await Injection("error");
        setError(errorResult);

        const blindBooleanResult = await Injection("boolean");
        setBlindBoolean(blindBooleanResult);

        const blindTimeResult = await Injection("time");
        setBlindTime(blindTimeResult);
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
                <Button variant={"faded"} className={`ml-2`} onPress={async () => await inject()}>Run</Button>
            </div>
            <Divider className={`my-4 w-[700px]`}/>
            {unionResult.length !== 0 ?
                <Table fullWidth={true} w>
                    <TableHeader>
                        <TableColumn>METHOD</TableColumn>
                        <TableColumn>DOMAIN</TableColumn>
                        <TableColumn>VULN?</TableColumn>
                        <TableColumn>PAYLOAD</TableColumn>
                    </TableHeader>
                    <TableBody>
                        <TableRow key={1}>
                            <TableCell><p className={`tracking-widest uppercase`}>UNION</p></TableCell>
                            <TableCell>{unionResult.domain}</TableCell>
                            <TableCell>
                                {
                                    unionResult.bool === "true" ?
                                        <Chip className="capitalize" color={"success"} size="sm" variant="flat">
                                            True
                                        </Chip>
                                        :
                                        <Chip className="capitalize" color={"danger"} size="sm" variant="flat">
                                            False
                                        </Chip>
                                }
                            </TableCell>
                            <TableCell className={`flex flex-row gap-2 items-center`}>
                                <p className={`bg-black/50 p-1 rounded-lg truncate w-[275px] items-center`}>{unionResult.result}</p>
                                <Tooltip content={"Copy Payload"} showArrow={true} closeDelay={0} delay={0}>
                                    <Button isIconOnly={true} size={"sm"} onPress={() => ClipboardSetText(unionResult.result)}>
                                        <Icon icon={"solar:copy-bold-duotone"} width={16} height={16} />
                                    </Button>
                                </Tooltip>
                            </TableCell>
                        </TableRow>
                        <TableRow key={2}>
                            <TableCell><p className={`tracking-widest uppercase`}>ERROR</p></TableCell>
                            <TableCell>{errorResult.domain}</TableCell>
                            <TableCell>
                                {
                                    errorResult.bool === "true" ?
                                        <Chip className="capitalize" color={"success"} size="sm" variant="flat">
                                            True
                                        </Chip>
                                        :
                                        <Chip className="capitalize" color={"danger"} size="sm" variant="flat">
                                            False
                                        </Chip>
                                }
                            </TableCell>
                            <TableCell className={`flex flex-row gap-2 items-center`}>
                                <p className={`bg-black/50 p-1 rounded-lg truncate w-[275px] items-center`}>{errorResult.result}</p>
                                <Tooltip content={"Copy Payload"} showArrow={true} closeDelay={0} delay={0}>
                                    <Button isIconOnly={true} size={"sm"} onPress={() => ClipboardSetText(errorResult.result)}>
                                        <Icon icon={"solar:copy-bold-duotone"} width={16} height={16} />
                                    </Button>
                                </Tooltip>
                            </TableCell>
                        </TableRow>
                        <TableRow key={3}>
                            <TableCell><p className={`tracking-widest uppercase`}>BLIND (BOOLEAN)</p></TableCell>
                            <TableCell>{blindBooleanResult.domain}</TableCell>
                            <TableCell>
                                {
                                    blindBooleanResult.bool === "true" ?
                                        <Chip className="capitalize" color={"success"} size="sm" variant="flat">
                                            True
                                        </Chip>
                                        :
                                        <Chip className="capitalize" color={"danger"} size="sm" variant="flat">
                                            False
                                        </Chip>
                                }
                            </TableCell>
                            <TableCell className={`flex flex-row gap-2 items-center`}>
                                <p className={`bg-black/50 p-1 rounded-lg truncate w-[275px] items-center`}>{blindBooleanResult.result}</p>
                                <Tooltip content={"Copy Payload"} showArrow={true} closeDelay={0} delay={0}>
                                    <Button isIconOnly={true} size={"sm"} onPress={() => ClipboardSetText(blindBooleanResult.result)}>
                                        <Icon icon={"solar:copy-bold-duotone"} width={16} height={16} />
                                    </Button>
                                </Tooltip>
                            </TableCell>
                        </TableRow>
                        <TableRow key={4}>
                            <TableCell><p className={`tracking-widest uppercase`}>BLIND (TIME)</p></TableCell>
                            <TableCell>{blindTimeResult.domain}</TableCell>
                            <TableCell>
                                {
                                    blindTimeResult.bool === "true" ?
                                        <Chip className="capitalize" color={"success"} size="sm" variant="flat">
                                            True
                                        </Chip>
                                        :
                                        <Chip className="capitalize" color={"danger"} size="sm" variant="flat">
                                            False
                                        </Chip>
                                }
                            </TableCell>
                            <TableCell className={`flex flex-row gap-2 items-center`}>
                                <p className={`bg-black/50 p-1 rounded-lg truncate w-[275px] items-center`}>{blindTimeResult.result}</p>
                                <Tooltip content={"Copy Payload"} showArrow={true} closeDelay={0} delay={0}>
                                    <Button isIconOnly={true} size={"sm"} onPress={() => ClipboardSetText(blindTimeResult.result)}>
                                        <Icon icon={"solar:copy-bold-duotone"} width={16} height={16} />
                                    </Button>
                                </Tooltip>
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
            <p>Tests if the site is vulnerable to SQL Injection by attempting various queries.</p>
            </footer>
        </div>
    )
}

function DisplayPayload({ result, method }) {
    return (
        <TableRow>
            <TableCell><p className={`tracking-widest uppercase`}>{method}</p></TableCell>
            <TableCell>{result.domain}</TableCell>
            <TableCell>
                {
                    result.bool === "true" ?
                        <Chip className="capitalize" color={"success"} size="sm" variant="flat">
                            True
                        </Chip>
                        :
                        <Chip className="capitalize" color={"danger"} size="sm" variant="flat">
                            False
                        </Chip>
                }
            </TableCell>
            <TableCell className={`flex flex-row gap-2 items-center`}>
                <p className={`bg-black/50 p-1 rounded-lg truncate w-[275px] items-center`}>{result.result}</p>
                <Tooltip content={"Copy Payload"} showArrow={true} closeDelay={0} delay={0}>
                    <Button isIconOnly={true} size={"sm"} onPress={() => ClipboardSetText(result.result)}>
                        <Icon icon={"solar:copy-bold-duotone"} width={16} height={16} />
                    </Button>
                </Tooltip>
            </TableCell>
        </TableRow>
    )
}