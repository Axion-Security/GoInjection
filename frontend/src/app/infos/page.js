"use client";

import {
    Button, Chip,
    Divider,
    Input, Pagination,
    Table,
    TableBody,
    TableCell,
    TableColumn,
    TableHeader,
    TableRow
} from "@nextui-org/react";
import React, {useEffect, useMemo, useState} from "react";
import {GetTarget, SetTarget, CheckWAF} from "../../../wailsjs/go/main/App";
import {Icon} from "@iconify/react";

export default function Page() {
    const [page, setPage] = useState(1);
    const [values, setValues] = useState([]);
    const rowsPerPage = 7;

    useEffect(() => {
        const fetchData = async () => {
            const data = [
                { option: 'URL', value: GetTarget("url")  },
                { option: 'DBMS', value: GetTarget("dbms") },
                { option: 'Syntax', value: GetTarget("syntax") },
                { option: 'Columns', value:  GetTarget("columns") },
                { option: 'Database', value: GetTarget("database") },
                { option: 'Union Injection', value: GetTarget("union") },
                { option: 'Error Injection', value: GetTarget("error") },
                { option: 'Boolean-based Injection', value: GetTarget("boolean") },
                { option: 'Time-based Injection', value: GetTarget("time") },
                { option: "Total Links sent for Union", value: GetTarget("unionLinks") },
                { option: "Total Links sent for Error", value: GetTarget("errorLinks") },
                { option: "Total Links sent for Boolean", value: GetTarget("booleanLinks") },
                { option: "Total Links sent for Time", value: GetTarget("timeLinks") },
                { option: "Total Links sent for WAF", value: GetTarget("wafLinks") },
                { option: "Total Links sent for Fingerprint", value: GetTarget("fingerprintLinks") },
            ];
            setValues(data);
        };

        fetchData();
    }, [])

    const pages = Math.ceil(values.length / rowsPerPage);

    const items = useMemo(() => {
        const start = (page - 1) * rowsPerPage;
        const end = start + rowsPerPage;

        return values.slice(start, end);
    }, [page, values]);

    return (
        <div className="flex flex-col justify-center items-center p-1">
            <Table fullWidth={true}>
                <TableHeader>
                    <TableColumn>OPTION</TableColumn>
                    <TableColumn>VALUE</TableColumn>
                </TableHeader>
                <TableBody>
                    {items.map((item, index) => (
                        <TableRow key={index}>
                            <TableCell>{item.option}</TableCell>
                                <TableCell>{item.value.length === 0 ? <p>No value.</p> : item.value}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>

            <Pagination
                isCompact
                showControls
                showShadow
                color="primary"
                page={page}
                total={pages}
                onChange={(newPage) => setPage(newPage)} // setPage is synchronous
                className="flex justify-center mt-1"
            />

            <footer className="bg-[#2c2d31] rounded p-2 flex flex-col justify-center items-center mt-4 text-sm text-gray-500 absolute bottom-5 w-[735px]">
                <p>Displays the collected information from the current target</p>
            </footer>
        </div>
    );
}