components {
  id: "script"
  component: "/background/bg_para_scroll2.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite2"
  type: "sprite"
  data: "tile_set: \"/background/background_para.atlas\"\n"
  "default_animation: \"bg_para_3\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 960.0
    y: 540.0
    z: -0.9
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
