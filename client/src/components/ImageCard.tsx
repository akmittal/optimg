import React, { ReactElement } from "react";
import {
  Card,
  CardActions,
  CardMedia,
  CardActionButton,
  CardActionButtons,
  CardPrimaryAction,
} from "@rmwc/card";
import { Typography } from "@rmwc/typography";

import "@material/card/dist/mdc.card.css";
import "@material/button/dist/mdc.button.css";
import "@material/icon-button/dist/mdc.icon-button.css";
import { Chip } from "rmwc";
import { Link } from "react-router-dom";

interface Props {
  full?: boolean;
  path?: string;
  varients?:number;
  imageURL: string;
  height: number;
  width: number;
  modified: string;
  size: number;
  name: string;
}

export default function ImageCard(props: Props): ReactElement {
  const { imageURL, height, width, modified, size, full, name, path, varients } = props;
  return (
    <Card>
      <CardPrimaryAction>
        <CardMedia
          sixteenByNine
          style={{
            backgroundImage: `url(${imageURL})`,
          }}
        />
        <div style={{ padding: "0 1rem 1rem 1rem" }}>
          <Typography use="headline6" tag="h2">
            {name}
          </Typography>

          <Typography
            className="text-left"
            use="body1"
            tag="div"
            theme="textSecondaryOnBackground"
          >
            Size: {Math.round(size / 1024)}KB
            <br />
            Dimensions: {width} x {height}
            <br />
            Modified: {new Date(modified).toLocaleString()}
            <br />
          </Typography>
        </div>
      </CardPrimaryAction>
      {!full && (
        <CardActions className="flex justify-between">
          <CardActionButtons>
            <Link to={`/gallery/image?path=${path}`}>
              <CardActionButton>Open</CardActionButton>
            </Link>
          </CardActionButtons>
          <div>
            <Chip selected label={`${varients} Varients`} />
          </div>
        </CardActions>
      )}
    </Card>
  );
}
