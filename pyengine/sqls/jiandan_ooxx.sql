/*
 Navicat Premium Data Transfer

 Source Server         : root
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : sys

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 19/02/2021 01:08:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for jiandan_ooxx
-- ----------------------------
DROP TABLE IF EXISTS `jiandan_ooxx`;
CREATE TABLE `jiandan_ooxx`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `date` date NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 335 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jiandan_ooxx
-- ----------------------------
INSERT INTO `jiandan_ooxx` VALUES (334, 'http://wx1.sinaimg.cn/large/00893JKXly1gnrz7eqb7jj30iq0xc0ui.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (333, 'http://wx3.sinaimg.cn/large/00893JKXly1gns061lvhij31400u0qv6.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (332, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gns0e14scxj30u0190u0x.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (331, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gns0jz4z1kj30jg0t64c3.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (330, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnrzc9c6vwj30u00u1dm5.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (329, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnrzi81fc9j30jg0t64eq.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (328, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnrzot7jksj30u019btr5.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (327, 'http://wx3.sinaimg.cn/large/0076BSS5ly1gnrzv0bv89j30u011iwmo.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (306, 'http://wx2.sinaimg.cn/large/0076BSS5ly1gnrxeg3m22j30pa0iy7a4.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (289, 'http://wx3.sinaimg.cn/large/0076BSS5ly1gnqoasonb0j31900u0qoy.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (279, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnqqe77yigj30u00k042i.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (280, 'http://wx3.sinaimg.cn/large/0076BSS5ly1gnqq8359y6j30to18g7fn.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (281, 'http://wx4.sinaimg.cn/large/00893JKXgy1gnqq1nxhi9j30ma0zkaen.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (282, 'http://wx2.sinaimg.cn/large/0076BSS5ly1gnqpow8x35j30sg0iymyr.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (283, 'http://wx3.sinaimg.cn/large/0076BSS5ly1gnqpjc9a91j30to18gk12.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (284, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnqpcxsbe2j30jg0t616b.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (326, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gns016jtvxj30u0190wlh.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (325, 'http://wx4.sinaimg.cn/large/00893JKXly1gns050n2exj30u0140qv5.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (324, 'http://wx2.sinaimg.cn/large/0076BSS5ly1gns076oaz0j30jg0t6any.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (323, 'http://wx2.sinaimg.cn/large/00869HDlly1gl2zxvmbsvg306k0gnka0.gif', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (322, 'http://wx3.sinaimg.cn/large/0076BSS5ly1gnrvc9g7pcj30iz0sgq63.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (321, 'http://wx1.sinaimg.cn/large/00893JKXly1gnrvc6htypg306e09lhdt.gif', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (320, 'http://wx2.sinaimg.cn/large/0076BSS5ly1gnrvhpq4iij31hd0u0dux.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (319, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnrvncnac2j30u0190woi.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (318, 'http://wx4.sinaimg.cn/large/00893JKXly1gnrvqtvkeyj30u01dd4qp.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (317, 'http://wx1.sinaimg.cn/large/00893JKXly1gnrvr0n2hbj30u0190kjl.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (316, 'http://wx3.sinaimg.cn/large/00893JKXly1gnrvrcnkwpj30u0190e83.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (315, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnrvt8rrfbj30rr0ihaov.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (314, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnrw6lapd8j30jg0sw7nk.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (313, 'http://wx2.sinaimg.cn/large/0076BSS5ly1gnrwd80y9kj30u018yke4.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (312, 'http://wx2.sinaimg.cn/large/0076BSS5ly1gnrwol3xmvj30u00u1ag6.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (311, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnrwub90ynj30u0190n3h.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (310, 'http://wx1.sinaimg.cn/large/00893JKXly1gnrwvvynhij30u01404np.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (309, 'http://wx2.sinaimg.cn/large/0076BSS5ly1gnrx18lvcxj30jg0t6wwp.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (308, 'http://wx1.sinaimg.cn/large/00893JKXly1gnrx2bf75oj30u0140x6q.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (307, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnrx7ncqzjj30sg0j0q4s.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (305, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnrxk3zdcsj31900u0ar6.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (304, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnrxw7xbjtj30iz0sgq5c.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (303, 'http://wx2.sinaimg.cn/large/0076BSS5ly1gnry2w67g5j312w0pzwry.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (302, 'http://wx3.sinaimg.cn/large/0076BSS5ly1gnry9ssfgjj30jg0t8tce.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (301, 'http://wx3.sinaimg.cn/large/0076BSS5ly1gnrygm4jxhj30u00u0jw6.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (300, 'http://wx3.sinaimg.cn/large/0076BSS5ly1gnrymlu7ygj31900u017e.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (299, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnrz5b46qdj30uk0kdh45.jpg', 'download', '2021-02-02');
INSERT INTO `jiandan_ooxx` VALUES (298, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnqsqgatkzj30u011jwg3.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (297, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnqt8hqvaaj30u0190kg5.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (296, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnqtfawhosj31gi0u01kx.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (295, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnqtl2bx6cj30f90lcdhd.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (294, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnqtqtjf50j30u01hcq8p.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (293, 'http://wx3.sinaimg.cn/large/0076BSS5ly1gnqtx12cknj31900u0wzh.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (292, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnqu3hi8qvj30u0190wsm.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (291, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnqua70v72j30tm18gq73.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (286, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnqotnquodj30u00u01ax.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (271, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnqrq9o69zj30u0190kdp.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (288, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnqogy1iklj30u00u0af7.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (285, 'http://wx3.sinaimg.cn/large/0076BSS5ly1gnqp0agg9ej30tn18gkd0.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (278, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnqqjq3tv0j30go0ci3yx.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (277, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnqquwonoxj30pa11x12r.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (276, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnqr0g5twjj30u01964hf.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (274, 'http://wx1.sinaimg.cn/large/0076BSS5ly1gnqrdq8rl4j31900u0b29.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (275, 'http://wx3.sinaimg.cn/large/0076BSS5ly1gnqr7927v0j30u0140q78.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (290, 'http://wx4.sinaimg.cn/large/0076BSS5ly1gnqugnfugij31900u0tfu.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (287, 'http://wx3.sinaimg.cn/large/0076BSS5ly1gnqoneje6rj30s10s177p.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (272, 'http://wx3.sinaimg.cn/large/0076BSS5ly1gnqrjx8ed0j30p70ixdl8.jpg', 'download', '2021-02-17');
INSERT INTO `jiandan_ooxx` VALUES (273, 'http://wx2.sinaimg.cn/large/00893JKXly1gnqre52ya7j30u01407wj.jpg', 'download', '2021-02-17');

SET FOREIGN_KEY_CHECKS = 1;
