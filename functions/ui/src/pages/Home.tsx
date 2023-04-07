import { Link } from "react-router-dom";
import "../styles/Home.css";
import { createRef, useEffect, useState } from "react";
import * as THREE from "three";
import SplineLoader from "@splinetool/loader";
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls";
import { Intro } from "../components/intro";

export default function Home() {
  const model = createRef<HTMLDivElement>();
  const [visible, setVisible] = useState(true);

  useEffect(() => {
    if (visible) {
      setTimeout(() => {
        setVisible(false);
      }, 4500);
    }
  }, [visible]);

  useEffect(() => {
    console.log("effect");
    const n = window.innerWidth > 1000 ? 1 : 1.5;
    // camera
    const camera = new THREE.OrthographicCamera(
      window.innerWidth / -n,
      window.innerWidth / n,
      window.innerHeight / n,
      window.innerHeight / -n,
      -50000,
      50000
    );
    camera.position.set(0, 1, -50);
    camera.quaternion.setFromEuler(new THREE.Euler(0, 0, 0));

    // scene
    const scene = new THREE.Scene();
    //const textureLoader = new THREE.TextureLoader();
    //scene.background = textureLoader.load("textures/background.png");

    // spline scene
    const loader = new SplineLoader();
    loader.load("scene.splinecode", (splineScene) => {
      scene.add(splineScene);
    });

    // renderer
    const renderer = new THREE.WebGLRenderer({ alpha: true });
    renderer.setSize(window.innerWidth, window.innerHeight);
    renderer.setAnimationLoop(animate);
    renderer.setClearColor(0x000000, 0);
    model.current?.appendChild(renderer.domElement);

    // scene settings
    renderer.shadowMap.enabled = true;
    renderer.shadowMap.type = THREE.PCFShadowMap;

    renderer.setClearAlpha(1);

    // orbit controls
    const controls = new OrbitControls(camera, model.current);
    controls.enableDamping = true;
    controls.dampingFactor = 0.125;
    controls.rotateSpeed = 0.4;
    controls.panSpeed = 0.4;
    controls.autoRotate = true;

    window.addEventListener("resize", onWindowResize);
    function onWindowResize() {
      camera.left = window.innerWidth / -2;
      camera.right = window.innerWidth / 2;
      camera.top = window.innerHeight / 2;
      camera.bottom = window.innerHeight / -2;
      camera.updateProjectionMatrix();
      renderer.setSize(window.innerWidth, window.innerHeight);
    }

    function animate() {
      renderer.render(scene, camera);
      controls.update();
    }
  }, [model]);

  return (
    <main className="main">
      <Intro active={visible}></Intro>
      <div ref={model} className="background"></div>
      <div className="actions">
        <Link to="/games/t-rex-game/gmp">
          <div className="play">
            <img className="playIcon" src="play_arrow.svg" alt="play" />
          </div>
        </Link>
        <Link to="/about" className="about">
          ABOUT
        </Link>
      </div>
    </main>
  );
}
