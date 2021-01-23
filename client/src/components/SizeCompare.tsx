import React, { ReactElement } from 'react';
import {UpOutlined, DownOutlined} from "@ant-design/icons"

interface Props {
    percentage: number
}

function SizeCompare({percentage}: Props): ReactElement {
    return (
        <strong className={percentage > 0 ? "text-red-500":"text-green-500"}>
          <div className="flex items-center">{percentage > 0 ? <UpOutlined /> : <DownOutlined />}{Math.abs(percentage)} {percentage > 0 ? " Bigger":" Smaller"}</div>
        </strong>
    )
}

export default SizeCompare;
