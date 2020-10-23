import React, { ReactElement, useEffect } from "react";
import { useQuery } from "react-query";
import { useLocation } from "react-router-dom";
import ImageCard from "./ImageCard";

interface Props {}

export default function GalleryImage(props: Props): ReactElement {
  const location = useLocation();
  const params = new URLSearchParams(location.search);
  const path = params.get("path") || "";
  const { isLoading, error, data } = useQuery("image", () =>
    fetch(`/img?path=${path}`).then((res) => res.json())
  );
  useEffect(() => {
    return () => {};
  }, [path]);
  if (isLoading) {
    return <div>Loading</div>;
  }
  if (error) {
    return <div>Errror</div>;
  }
  // const date = new Date()
  return (
    <div>
      <ImageCard
        height={data["Main"]["Height"]}
        width={data["Main"]["Width"]}
        name={data["Main"]["Path"]}
        imageURL={`/images/src${data["Main"]["Path"]}`}
        modified={data["Main"]["Modified"]}
        size={data["Main"]["Size"]}
        full
      />
      <h3 className="text-4xl text-left">Varients</h3>
      <div className="flex justify-evenly">{data["Varients"].map((varient: any) => {
        return (
          <ImageCard
            height={varient["Height"]}
            width={varient["Width"]}
            name={varient["Path"]}
            imageURL={`/images/target${varient["Path"]}`}
            modified={varient["Modified"]}
            size={varient["Size"]}
            full
          />
        );
      })}</div>
    </div>
  );
}
