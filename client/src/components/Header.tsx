import React, { ReactElement } from "react";
import { Typography } from "@rmwc/typography";

interface Props {
  text: string;
}

export default function Header(props: Props): ReactElement {
  return (
    <Typography use="headline3" className="text-left m-6 my-10 font-light">
      {props.text}
    </Typography>
  );
}
